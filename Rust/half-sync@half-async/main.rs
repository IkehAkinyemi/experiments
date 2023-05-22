use std::sync::{Arc, Mutex};
use std::thread;
use tokio::task;

fn main() {
    // Shared state between threads
    let shared_state = Arc::new(Mutex::new(0));

    // Synchronous part of the system
    for _ in 0..5 {
        let shared_state = Arc::clone(&shared_state);
        thread::spawn(move || {
            let mut data = shared_state.lock().unwrap();
            // Simulate some synchronous computation
            *data += 1;
            println!("Synchronous data: {}", *data);
        });
    }

    // Asynchronous part of the system
    let shared_state_async = Arc::clone(&shared_state);
    tokio::runtime::Runtime::new().unwrap().block_on(async {
        let mut handles = Vec::new();
        for _ in 0..5 {
            let shared_state = Arc::clone(&shared_state_async);
            let handle = task::spawn(async move {
                let mut data = shared_state.lock().unwrap().await;
                // Simulate some asynchronous computation
                *data += 1;
                println!("Asynchronous data: {}", *data);
            });
            handles.push(handle);
        }
        for handle in handles {
            handle.await.unwrap();
        }
    });
}
