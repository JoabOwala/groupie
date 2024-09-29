# Groupie Tracker

## Overview

Groupie Trackers is a web application designed to fetch and manipulate data from a specified APIs. The web app is intuitive showcases information about bands and artists. This project involves working with multiple data entities and ensuring a seamless client-server interaction.

## API Structure

The application interacts with an API divided into four main parts:

1. **Artists**: Contains information about bands and artists, including:
   - Name(s)
   - Image
   - Year of activity
   - Date of the first album
   - Members

2. **Locations**: Lists the last and upcoming concert locations.

3. **Dates**: Provides information on the last and upcoming concert dates.

4. **Relation**: Links artists, dates, and locations together.

## Features

- User-friendly website displaying band information.
- Data visualizations using cards.
- Client-server communication to trigger actions and retrieve information.

## Event Handling

The application implements a client-server model where user actions (e.g., button clicks) trigger requests to the server, which respond with the relevant data. 

## Technologies Used

- **Go**: Backend server and template rendering.
- **HTML/CSS**: Frontend design and user interface.

## Getting Started

To get started with this project, follow these instructions:

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installationn

1. **Clone the Repository:**

```bash
git clone https://learn.zone01kisumu.ke/git/cowalla/groupie-tracker
```
```bash
cd groupie-tracker
```
2. **Run the Application:**

    Start the server with:

    ```bash
    go run .
    ```

    By default, the server will run on `http://localhost:8000`.

3. **Access the Application:**

    Open your web browser and navigate to `http://localhost:8000` to access the Groupie-tracker Web application.

## Example
![Example](assets/groupie-tracker.gif)

## Contributing

Contributions are welcome! If you have suggestions or improvements, please fork the repository and submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

## Authors

- [Joab Owala](https://github.com/jowala)
- [Cheril Owala](https://github.com/cowalla)