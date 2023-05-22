use std::sync::mpsc;
use std::thread;

// Define the ActiveObject struct
struct ActiveObject {
    sender: mpsc::Sender<Box<dyn FnOnce() + Send>>,
}

impl ActiveObject {
    fn new() -> ActiveObject {
        let (sender, receiver): (mpsc::Sender<Box<dyn FnOnce() + Send>>, mpsc::Receiver<Box<dyn FnOnce() + Send>>) = mpsc::channel();

        // Spawn a worker thread
        thread::spawn(move || {
            while let Ok(task) = receiver.recv() {
                task();
            }
        });

        ActiveObject { sender }
    }

    fn do_task<F>(&self, task: F)
    where
        F: FnOnce() + Send + 'static,
    {
        self.sender.send(Box::new(task)).unwrap();
    }
}

// Usage example
fn main() {
    let active_object = ActiveObject::new();

    // Perform tasks asynchronously through the active object
    active_object.do_task(|| {
        println!("Task 1: Executing...");
    });

    active_object.do_task(|| {
        println!("Task 2: Executing...");
    });

    // Wait for tasks to complete
    thread::sleep(std::time::Duration::from_secs(1));
}
