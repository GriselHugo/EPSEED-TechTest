import { useRef } from "react";
import { Link } from "react-router-dom";
import { FaBars, FaTimes } from 'react-icons/fa';

import Button from 'react-bootstrap/Button';

import './Navbar.css';

function Navbar() {
    const navRef = useRef(null);

    const toggleNav = () => {
        navRef.current.classList.toggle("show");
    }

    const logOut = () => {
        localStorage.setItem("isLogged", 0)
        localStorage.setItem("currentUserId", -1);
        window.location.reload();
    }

    return (
        <header id="navHeader">
            <h3>EPSEED TECH TEST</h3>

            <nav id="navBar" ref={navRef}>
                <Link to="/">Home</Link>
                <Link to="/note">Note</Link>
                <button className="nav-btn nav-close-btn" onClick={toggleNav}>
                    <FaTimes />
                </button>
                <Button variant="danger" onClick={logOut}>Logout</Button>
            </nav>
            <button id="navBurger" className="nav-btn" onClick={toggleNav}>
                <FaBars />
            </button>
        </header>
    );
}

export default Navbar;
