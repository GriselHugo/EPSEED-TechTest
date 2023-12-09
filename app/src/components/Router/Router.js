import { Routes, Route } from 'react-router-dom';

import Home from '../Home/Home';
import Note from '../Note/Note';

function Router() {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/note" element={<Note />} />
        </Routes>
    );
}

export default Router;
