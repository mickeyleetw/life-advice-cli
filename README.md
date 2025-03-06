# Life Advice CLI

This is a command-line interface (CLI) application that provides life advice in the form of quotes and jokes. It fetches data from external APIs and maintains a history of fetched advice.

## Features

*   **Fetch Jokes:** Retrieves jokes from the `https://icanhazdadjoke.com/` API.
*   **Fetch Quotes:** Retrieves quotes from the `http://api.quotable.io/random` API.
*   **Advice History:** Keeps track of the fetched advice, allowing users to view their recent history.
*   **Multiple Advice Types:**  Allows fetching multiple types of advice in a single command.
*   **Error Handling:** Gracefully handles API errors and unknown advice types.
*   **Concurrency:** Uses goroutines to fetch advice concurrently, improving performance.

## Installation
To run this CLI, you need to have Go installed on your system.
1.  Clone this repository:
    ```bash
    git clone https://github.com/mickeyleetw/life-advice-cli.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd life-advice-cli
    ```
3. Build the project
   ```bash
    go build -o advice main/main.go
   ```
4. Run the project
   ```bash
    ./advice --advice="quote, joke" --history=true
   ```

## Usage

The CLI supports the following flags:

*   `--advice`: Specifies the type(s) of advice to fetch.  Multiple types can be separated by commas (e.g., `quote,joke`).  Supported types are:
    *   `joke`
    *   `quote`
*   `--history`: Displays the history of fetched advice.

**Examples:**

1.  **Fetch a joke:**

    ```bash
    ./advice -advice joke
    ```

2.  **Fetch a quote:**

    ```bash
    ./advice --advice quote
    ```

3.  **Fetch both a joke and a quote:**

    ```bash
    ./advice --advice "joke, quote"
    ```

4.  **Show the advice history:**

    ```bash
    ./advice --history
    ```


## Code Structure

The project is organized into the following packages:

*   `main`: Contains the main entry point of the application. It handles command-line flags, initializes the necessary components, and manages the fetching and display of advice.
*   `core`:
    *   `fetcher.go`: Defines the `AdviceFetcher` struct and related methods for fetching advice from external APIs. It includes error handling and response parsing.  It defines an `AdviceFetcherInterface` for flexibility.
    *   `history.go`:  Implements a circular buffer (`History` struct) to store the history of fetched advice. It provides methods to add new records and retrieve the history.

## Concurrency

The application uses goroutines and channels to fetch advice from multiple APIs concurrently. This approach improves the responsiveness of the CLI when multiple advice types are requested.  A `sync.WaitGroup` is used to ensure all fetch operations complete before the program exits. Error handling is done via a dedicated error channel.

## Error Handling

*   **API Errors:** If an API request fails (e.g., network error, API unavailable), an error message is printed to the console.
*   **Unknown Advice Types:** If the user provides an unsupported advice type, a warning message is displayed.
* **Timeout:** The http client has a timeout of 5 seconds.

## History Implementation
The history is implemented as a circular buffer. The size of the buffer is fixed at the time of creation. When the buffer is full, the oldest record is overwritten by the newest one.
