import React from 'react';
import { useBooks } from '../context/BookContext';
import AddBook from './AddBook';

const BookList = () => {
    const { books, loading, error, addBook } = useBooks();

    if (loading) return <p>Loading...</p>;
    if (error) return <p>{error}</p>;

    return (
        <div>
            <h2>Books List</h2>
            <AddBook onAdd={addBook} />
            <ul>
                {books.map(book => (
                    <li key={book.id}>{book.title} by {book.author}</li>
                ))}
            </ul>
        </div>
    );
};

export default BookList;
