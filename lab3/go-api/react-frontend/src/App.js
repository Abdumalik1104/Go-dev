import React, { useState } from 'react';
import BookList from './components/BookList';
import { BookProvider } from './context/BookContext';
import Login from './components/Login';
import ProtectedComponent from './components/ProtectedComponent'; 

function App() {
    const [token, setToken] = useState(localStorage.getItem('token'));

    const saveToken = (userToken) => {
        localStorage.setItem('token', userToken); 
        setToken(userToken);
    };

    const logout = () => {
        localStorage.removeItem('token');
        setToken(null);
    };

    return (
        <BookProvider>
            <div className="App">
                <h1>My Book Collection</h1>
                {!token ? (
                    <Login setToken={saveToken} />
                ) : (
                    <>
                        <button onClick={logout}>Logout</button>
                        <BookList />
                        <ProtectedComponent token={token} /> {}
                    </>
                )}
            </div>
        </BookProvider>
    );
}

export default App;
