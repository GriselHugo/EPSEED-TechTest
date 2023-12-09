import React from "react";

function Home() {
    const currentUserId = parseInt(localStorage.getItem("currentUserId") || -1);

    return (
        <div className="Home">
            <h1>Home</h1>
            <p>Home page content</p>
            <p>Curren User id: {currentUserId}</p>
        </div>
    );
}

export default Home;
