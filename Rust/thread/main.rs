use std::thread;

fn main() {
    let handle = thread::spawn(|| {
        // code to be executed in the new thread
    });
}