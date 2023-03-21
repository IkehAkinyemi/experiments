use std::sync::{Arc, Mutex, MutexGuard};
use std::thread;

fn main() {
    let shared_data = Arc::new(Mutex::new(vec![1, 2, 3]));
    let mut handles = Vec::new();

    // Spawn two threads that will access the shared data
    for i in 0..2 {
        let shared_data = shared_data.clone(); // Clone the Arc to move into the thread
        let handle = thread::spawn(move || {
            let mut data: MutexGuard<Vec<i32>> = match shared_data.lock() {
                Ok(guard) => guard,
                Err(poisoned) => {
                    // Handle mutex poisoning
                    let guard = poisoned.into_inner();
                    println!("Thread {} recovered from mutex poisoning: {:?}", i, *guard);
                    guard
                }
            };

            // Use the data
            println!("Thread {}: {:?}", i, *data);
            data.push(i);
        });
        handles.push(handle);
    }

    // Wait for the threads to finish
    for handle in handles {
        handle.join().unwrap();
    }
}
