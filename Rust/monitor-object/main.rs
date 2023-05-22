use std::sync::{Arc, Mutex};
use std::thread;

// Define a shared resource protected by a Mutex
struct Counter {
    count: u32,
    mutex: Mutex<()>,
}

impl Counter {
    fn new() -> Self {
        Counter {
            count: 0,
            mutex: Mutex::new(()),
        }
    }

    fn increment(&self) {
        let _lock = self.mutex.lock().unwrap();
        self.count += 1;
    }

    fn get_count(&self) -> u32 {
        self.count
    }
}

fn main() {
    // Create a shared instance of Counter using Arc
    let counter = Arc::new(Counter::new());

    // Spawn multiple threads to increment the counter concurrently
    let mut handles = vec![];
    for _ in 0..5 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            for _ in 0..1000 {
                counter.increment();
            }
        });
        handles.push(handle);
    }

    // Wait for all threads to complete
    for handle in handles {
        handle.join().unwrap();
    }

    // Get the final count
    let final_count = counter.get_count();
    println!("Final count: {}", final_count);
}
