use std::cell::RefCell;

thread_local! {
    static THREAD_DATA: RefCell<u32> = RefCell::new(0);
}

fn main() {
    THREAD_DATA.with(|data| {
        *data.borrow_mut() += 1;
        println!("Thread-specific data: {}", *data.borrow());
    });

    // Spawn a new thread
    let handle = std::thread::spawn(|| {
        THREAD_DATA.with(|data| {
            *data.borrow_mut() += 2;
            println!("Thread-specific data in the spawned thread: {}", *data.borrow());
        });
    });

    handle.join().unwrap();

    THREAD_DATA.with(|data| {
        println!("Thread-specific data in the main thread: {}", *data.borrow());
    });
}
