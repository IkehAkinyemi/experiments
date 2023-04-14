use tracing_mutex::stdsync::TracingMutex;

fn main() {
    // Create a TracingMutex that guards access to a string
    let data = TracingMutex::new(42);
    
    // Spawn two threads that will access the shared date
    for i in 0..2 {
        let data = &data;
        std::thread::spawn(|| {
            // Lock the mutex and print the contents of the string
            let mut guard = data.lock().unwrap();
            println!("Thread {} says: {}", i, *guard);
            
            // Modify the string and release the mutex
            *guard = 43;
        });
    }
}
