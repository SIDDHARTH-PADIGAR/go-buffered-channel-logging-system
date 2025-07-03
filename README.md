###  **Use Case for Buffered Channel in Go**

**Problem:** You're producing data faster than you consume it, but you don't want the producer to block every time it sends a message.

###  Real-world Scenario

**Logging System** – where the app logs messages rapidly, but writing to disk or external system (like a database or file) is slow. You want the app to **continue execution without waiting** for the log to write.

---

###  Concept

* **Buffered Channel** acts as a queue.
* Producer sends logs → buffered channel holds them temporarily → consumer processes them at its own pace.
* If the buffer fills up, then and only then will the producer block.

---

###  Output (approximate timing):

```
>>> Attempting to send: Log entry #1
✓✓✓ Sent: Log entry #1
>>> Attempting to send: Log entry #2
✓✓✓ Sent: Log entry #2
>>> Attempting to send: Log entry #3
✓✓✓ Sent: Log entry #3
>>> Attempting to send: Log entry #4
✓✓✓ Sent: Log entry #4
>>> Attempting to send: Log entry #5
Logged: Log entry #1
✓✓✓ Sent: Log entry #5
```

---

###  Why This Works Well

* First 3 logs are pushed into the buffer **instantly**, because the channel has a capacity of 3.
* The `Logger` goroutine starts draining the channel with a 500ms delay per message.
* 4th and 5th log entries block at the `logChannel <- logMessage` line **only when the buffer is full**, and **only resume after the consumer frees up space** by reading.
* The `fmt.Printf("✓✓✓ Sent: ...")` only runs **after the message has been successfully sent**, making the blocking behavior crystal clear.
* This model is perfect for high-throughput systems like:

  * Logging systems
  * Event queues
  * Metric ingestion
  * Asynchronous job dispatchers
    where you want the **producer to stay fast** without constantly waiting for the consumer.

---

###  When to Use Buffered Channels

* When producers generate data faster than consumers can process.
* When you want to **reduce blocking** and still maintain order.
* Good for:

  * Logging systems
  * Job queues
  * Metrics pipelines
  * Data ingestion from multiple sensors, events, etc.
