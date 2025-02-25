
### Task: Library Management System

#### Objective:
Create a program that manages a library's books. The program will allow users to add books, display them, and sort them by the year of publication.

#### Requirements:

1. **Data Structure**:
   - Store each book's information (title, author, year of publication) in a simple data structure such as a slice (or array).
   - Do **not** use structs for this task. Instead, use arrays or slices to hold the book information.

2. **Menu Options**:
   The program should provide a simple menu with the following options:
   - **Add a book**: The user will be asked to input the title, author, and year of publication for a new book. This book will be added to the list of books.
   - **Display books**: The program will display all the books in the library with their details (title, author, year).
   - **Sort books by year of publication**: The program will sort the books in ascending order based on the year of publication.
   - **Exit**: Exit the program.

3. **Functions**:
   Write functions to handle each of the following tasks:
   - Adding a book to the list.
   - Displaying all books in the library.
   - Sorting the books by year of publication.

4. **Input/Output**:
   - When adding a book, prompt the user to enter the book's title, author, and year of publication.
   - After displaying the books, show the list with the format:
     ```
     Title: <title>, Author: <author>, Year of Publication: <year>
     ```
   - After sorting the books, display a message saying: `Books have been sorted by year of publication`.

#### Example of the Program Flow:

```
Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 1
Enter the book title: The Catcher in the Rye
Enter the book author: J.D. Salinger
Enter the year of publication: 1951

Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 1
Enter the book title: To Kill a Mockingbird
Enter the book author: Harper Lee
Enter the year of publication: 1960

Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 2
Title: The Catcher in the Rye, Author: J.D. Salinger, Year of Publication: 1951
Title: To Kill a Mockingbird, Author: Harper Lee, Year of Publication: 1960

Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 3
Books have been sorted by year of publication.

Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 2
Title: The Catcher in the Rye, Author: J.D. Salinger, Year of Publication: 1951
Title: To Kill a Mockingbird, Author: Harper Lee, Year of Publication: 1960

Library Management Program Menu:
1. Add a book
2. Display books
3. Sort books by year of publication
4. Exit
Please enter your choice: 4
Exiting the program.
```

#### Tips:
- Use arrays or slices to store the book data.
- Ensure that when sorting the books, you are comparing the year of publication and sorting them correctly.
- Handle user input carefully to avoid errors like entering non-numeric values for the year.
