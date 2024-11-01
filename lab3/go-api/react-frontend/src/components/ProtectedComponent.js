import React, { useEffect, useState } from 'react';
import axios from 'axios';

function ProtectedComponent({ token }) {
    const [data, setData] = useState(null);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/user', {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setData(response.data);
            } catch (error) {
                console.error("Error fetching user data:", error);
                setData({ message: "Access denied" });
            }
        };

        const fetchAdminData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/admin', {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                setData(response.data);
            } catch (error) {
                console.error("Error fetching admin data:", error);
                setData({ message: "Access denied" });
            }
        };

        fetchUserData();
        fetchAdminData();
    }, [token]);

    return (
        <div>
            <h2>Protected Data:</h2>
            {data ? <p>{JSON.stringify(data)}</p> : <p>Loading...</p>}
        </div>
    );
}

export default ProtectedComponent;
