import React, { useEffect, useState } from "react";
import goServer from "../../api/go-server";

import "./Note.css";

const NoteModifier = () => {
    return (
        <div className="NoteModifier">
            <h2>Modifier une note</h2>
            <form>
                <label>
                    Titre :
                    <br />
                    <input type="text" name="title" />
                </label>
                <br />
                <label>
                    Contenu :
                    <br />
                    <textarea
                        type="text"
                        name="content"
                        style={{ height: '200px', resize: 'none' }}
                    />
                </label>
                <br />
                <button type="submit">Enregistrer la note</button>
            </form>
        </div>
    );
};

function Note() {
    const [newNote, setNewNote] = useState({ title: '', content: '' });
    const [notes, setNotes] = useState([]);

    useEffect(() => {
            console.log('Current user id :', parseInt(localStorage.getItem("currentUserId")));
            goServer.getNotes(parseInt(localStorage.getItem("currentUserId"))).then((response) => {
            console.log("Get notes response: ", response);
            setNotes(response);
        }).catch((error) => {
            console.log("Error : " + error);
        });
    }, []);



    const handleChange = (e) => {
      const { name, value } = e.target;
      setNewNote((prevNote) => ({ ...prevNote, [name]: value }));
    };

    const handleSubmit = (e) => {
            if (newNote.title === "" || newNote.content === "") {
                return;
            }
            e.preventDefault();
            console.log('Nouvelle note :', newNote);
            goServer.createNote(parseInt(localStorage.getItem("currentUserId")), newNote.title, newNote.content).then((response) => {
            console.log("Create note response: ", response);
            const noteToAdd = {
                ID: response.ID,
                Title: newNote.title,
                Content: newNote.content
            };
            setNotes((prevNotes) => [...prevNotes, noteToAdd]);
        }).catch((error) => {
            console.log("Error : " + error);
        });
        setNewNote({ title: '', content: '' });
    };

    return (
        <div className="NoteContainer">

        <div className="NoteForm">
            <h2>Ajouter une note</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    Titre :
                    <br />
                    <input
                        type="text"
                        name="title"
                        value={newNote.title}
                        onChange={handleChange}
                    />
                </label>
                <br />
                <label>
                    Contenu :
                    <br />
                    <textarea
                        type="text"
                        name="content"
                        value={newNote.content}
                        style={{ height: '200px', resize: 'none' }}
                        onChange={handleChange}
                    />
                </label>
                <br />
                <button type="submit" className="submitNote">Enregistrer la note</button>
            </form>
        </div>

        <div className="NoteList">
            <h2>Liste des notes</h2>
            { notes.length === 0 ? <p>Aucune note à afficher </p> :
                <ul>
                    {notes.map((note) => (
                        <li key={note.ID}>
                            <h3>{note.Title}</h3><br />
                            <p>{note.Content}</p>
                        </li>
                    ))}
                </ul>
            }
        </div>

            <NoteModifier />
        </div>
    );
}

export default Note;
