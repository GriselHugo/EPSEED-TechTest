import React, { useEffect, useState } from "react";
import goServer from "../../api/go-server";

import "./Note.css";

function Note() {
    const [newNote, setNewNote] = useState({ title: '', content: '' });
    const [notes, setNotes] = useState([]);
    const [showNoteModifier, setShowNoteModifier] = useState(false);
    const [noteToModify, setNoteToModify] = useState({ title: '', content: '', id: '' });

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

    const handleModify = (e) => {
        const { name, value } = e.target;
        setNoteToModify((prevNote) => ({ ...prevNote, [name]: value }));
    }

    const handleSubmit = (e) => {
            if (newNote.title === "" || newNote.content === "") {
                return;
            }
            e.preventDefault();
            console.log('Nouvelle note :', newNote);
            goServer.createNote(parseInt(localStorage.getItem("currentUserId")), newNote.title, newNote.content).then((response) => {
            console.log("Create note response: ", response);
            const noteToAdd = {
                ID: response.id,
                Title: newNote.title,
                Content: newNote.content
            };
            setNotes((prevNotes) => [...prevNotes, noteToAdd]);
        }).catch((error) => {
            console.log("Error : " + error);
        });
        setNewNote({ title: '', content: '' });
    };

    const handleSubmitModification = (e) => {
        if (noteToModify.title === "" || noteToModify.content === "") {
            return;
        }
        e.preventDefault();
        console.log('Note à modifier :', noteToModify);
        goServer.updateNoteForUser(parseInt(localStorage.getItem("currentUserId")), noteToModify.id, noteToModify.title, noteToModify.content).then((response) => {
            console.log("Update note response: ", response.id);
            const noteToAdd = {
                ID: response.id,
                Title: noteToModify.title,
                Content: noteToModify.content
            };
            const index = notes.findIndex((note) => note.ID === noteToModify.id);
            const newNotes = [...notes];
            newNotes[index] = noteToAdd;
            setNotes(newNotes);
            setShowNoteModifier(false);
        }).catch((error) => {
            console.log("Error : " + error);
        });
        setNoteToModify({ title: '', content: '', id: '' });
    }

    const deleteNote = () => {
        console.log('Note à supprimer :', noteToModify);
        goServer.deleteNoteForUser(parseInt(localStorage.getItem("currentUserId")), noteToModify.id).then((response) => {
            console.log("Delete note response: ", response);
            const index = notes.findIndex((note) => note.ID === noteToModify.id);
            const newNotes = [...notes];
            newNotes.splice(index, 1);
            setNotes(newNotes);
            setShowNoteModifier(false);
        }).catch((error) => {
            console.log("Error : " + error);
        });
        setNoteToModify({ title: '', content: '', id: '' });
    }

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
                <div className="NoteListHeader">
                    {notes.map((note) => (
                        <div className="NoteListItem" key={note.ID} onClick={
                            () => {
                                console.log('Note cliquée :', note.ID)
                                setNoteToModify({ title: note.Title, content: note.Content, id: note.ID });
                                setShowNoteModifier(true);
                            }
                        } >
                            <p><h3>{note.Title}</h3></p>
                            <p>{note.Content}</p>
                        </div>
                    ))}
                </div>
            }
        </div>

        <div className="NoteModifier">
            <h2>Modifier une note</h2>
            { showNoteModifier ? (
            <form onSubmit={handleSubmitModification}>
            <label>
                Titre :
                <br />
                <input
                    type="text"
                    name="title"
                    value={noteToModify.title}
                    onChange={handleModify}
                />
            </label>
            <br />
            <label>
                Contenu :
                <br />
                <textarea
                    type="text"
                    name="content"
                    value={noteToModify.content}
                    style={{ height: '200px', resize: 'none' }}
                    onChange={handleModify}
                />
            </label>
            <br />
            <button type="submit" className="submitNote">Modifier la note</button>
            <button type="button" className="removeNote" onClick={() => deleteNote()}>Supprimer la note</button>
            <button type="button" className="submitNote" onClick={() => {
                setShowNoteModifier(false);
                setNewNote({ title: '', content: '', id: '' });
            }}>Annuler</button>
        </form>
            ) : <p>Choisissez une note à modifier</p>}
        </div>

        </div>
    );
}

export default Note;
