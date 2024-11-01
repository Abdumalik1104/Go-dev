import React, { useState } from 'react';
import axios from 'axios';
import { useBooks } from '../context/BookContext';

const AddBook = () => {
    const [title, setTitle] = useState('');
    const [author, setAuthor] = useState('');
    const { addBook } = useBooks();

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.post('http://localhost:8080/books', { title, author })
            .then(response => {
                addBook(response.data);
                setTitle('');
                setAuthor('');
            })
            .catch(error => console.error("There was an error creating the book!", error));
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="Title"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                required
            />
            <input
                type="text"
                placeholder="Author"
                value={author}
                onChange={(e) => setAuthor(e.target.value)}
                required
            />
            <button type="submit">Add Book</button>
        </form>
    );
};

export default AddBook;
