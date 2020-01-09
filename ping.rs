use std::io;
use std::io::prelude::*;
use std::net::{TcpListener, TcpStream};
use std::{thread, time};

fn main() -> io::Result<()> {
    let listener = TcpListener::bind("127.0.0.1:8080")?;

    thread::spawn(move || -> io::Result<()> {
        let mut stream = TcpStream::connect("127.0.0.1:8080")?;
        loop {
            thread::sleep(time::Duration::from_secs(2));
            stream.write("ping\n".as_bytes())?;
        }
    });

    for stream in listener.incoming() {
        thread::spawn(move || -> io::Result<()> {
            for byte in (&stream?).bytes() {
                let v = vec![byte?];
                print!("{}", String::from_utf8(v).unwrap_or_default());
            }
            Ok(())
        });
    }

    Ok(())
}
