import React, { useState, useEffect } from "react";
import goServer from "../../api/go-server";

function Home() {
    // const currentUserId = parseInt(localStorage.getItem("currentUserId") || -1);

    const [users, setUsers] = useState([]);

    useEffect(() => {
        goServer.getUsers().then((response) => {
            setUsers(response);
        }).catch((error) => {
            console.log("Error : " + error);
        });
    }, []);

    return (
        <div className="Home">
            <h1>Home</h1>
            <h2>List of users :</h2>
            <ul>
                { users.map((user) => (
                    <li key={user.ID}>{user.Username} - {user.Email}</li>
                ))
                }
            </ul>
        </div>
    );
}

export default Home;
