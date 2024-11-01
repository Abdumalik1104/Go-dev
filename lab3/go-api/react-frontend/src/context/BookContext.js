import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from 'axios';

const BookContext = createContext();

export const useBooks = () => useContext(BookContext);

export const BookProvider = ({ children }) => {
    const [books, setBooks] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        axios.get('http://localhost:8080/books')
            .then(response => {
                setBooks(response.data);
                setLoading(false);
            })
            .catch(error => {
                setError("There was an error fetching books!");
                setLoading(false);
            });
    }, []);

    const addBook = (newBook) => {
        setBooks([...books, newBook]);
    };

    return (
        <BookContext.Provider value={{ books, loading, error, addBook }}>
            {children}
        </BookContext.Provider>
    );
};
