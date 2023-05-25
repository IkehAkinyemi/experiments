use std::sync::{Arc, Mutex};
use std::sync::mpsc::{Sender, Receiver};
use std::thread;

struct Job {
    id: u32,
    // Additional fields related to the job
}

struct Worker {
    id: u32,
    thread_handle: Option<thread::JoinHandle<()>>,
}

impl Worker {
    fn new(id: u32, receiver: Arc<Mutex<Receiver<Job>>>) -> Worker {
        let thread_handle = thread::spawn(move || {
            loop {
                let job = {
                    let receiver = receiver.lock().unwrap();
                    receiver.recv().unwrap()
                };

                println!("Worker {} got job {}", id, job.id);
                // Perform the job's task here

                // Simulating some work
                thread::sleep(std::time::Duration::from_secs(1));
            }
        });

        Worker {
            id,
            thread_handle: Some(thread_handle),
        }
    }
}

impl Drop for Worker {
    fn drop(&mut self) {
        if let Some(handle) = self.thread_handle.take() {
            handle.join().unwrap();
        }
    }
}

struct ThreadPool {
    workers: Vec<Worker>,
    sender: Sender<Job>,
}

impl ThreadPool {
    fn new(size: u32) -> ThreadPool {
        assert!(size > 0);

        let (sender, receiver) = std::sync::mpsc::channel();
        let receiver = Arc::new(Mutex::new(receiver));

        let mut workers = Vec::with_capacity(size as usize);
        for id in 0..size {
            workers.push(Worker::new(id, Arc::clone(&receiver)));
        }

        ThreadPool {
            workers,
            sender,
        }
    }

    fn execute(&self, job: Job) {
        self.sender.send(job).unwrap();
    }
}

fn main() {
    let thread_pool = ThreadPool::new(4);

    for job_id in 0..10 {
        let job = Job { id: job_id };
        thread_pool.execute(job);
    }

    // Simulating work in the main thread
    thread::sleep(std::time::Duration::from_secs(5));
}
