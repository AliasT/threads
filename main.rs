
use std::thread;
use std::sync::mpsc;
use std::time::Duration;

fn main() {
  let (tx, rx) = mpsc::channel();
  let mut handles = Vec::new();
  for i in 0..3 {
    let sender = tx.clone();
    let handle = thread::spawn(move || {
      thread::sleep(Duration::from_secs(1));
      sender.send(i).unwrap();
    });
    handles.push(handle);
  }

  for h in handles {
    // 在等待的过程中, 其他进程可能已完成，所以channel中的数据输出不固定。
    h.join().unwrap();
    println!("{}", rx.recv().unwrap());
  }
}
