## Tests
In order to run tests for Anansi profiler, you can follow the steps below:

1. Create a ClickHouse database instance
2. Run the `create-and-insert.sql` script to create the tables and insert some data (be aware that this dataset is big and can take a while to insert)
3. Capture the logs from the ClickHouse instance using a tail command and save them to a file that can be used as input for the Anansi profiler
4. Run the queries you want while capturing the logs (You can use the `queries.sql` script in this folder as an example)
