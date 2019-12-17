const { Worker } = require('worker_threads');

for (let i = 0; i < 3; i++) {
  const worker = new Worker(`
    const { parentPort } = require('worker_threads');
    setTimeout(() => parentPort.postMessage(${i}))
  `, { eval: true });
  worker.on('message', message => console.log(message));
}
