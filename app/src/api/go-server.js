import axios from 'axios';

class GoServer {
  constructor() {
    this.client = axios.create({
      baseURL: 'http://localhost:8080',
    });
  }

    async signUp( username, email, password ) {
        try {
            const body = {
                username: username,
                email: email,
                password: password
            };

            const response = await this.client.post('/signup', body);

            // console.log("Sign up response: ", response );

            return response.data;
        } catch (error) {
            console.error("Error signing up: ", error.message);
            throw error;
        }
    }

    async logIn( email, password ) {
        try {
            const body = {
                email: email,
                password: password
            };

            const response = await this.client.post('/login', body);

            // console.log("Log in response: ", response );

            return response.data;
        } catch (error) {
            console.error("Error signing in: ", error.message);
            throw error;
        }
    }

    async getUsers() {
        try {
            const response = await this.client.get('/getall');

            // console.log("Get users response: ", response.data );

            return response.data;
        } catch (error) {
            console.error("Error getting users: ", error.message);
            throw error;
        }
    }

    async createNote( userId, title, content ) {
        try {
            const body = {
                user_id: userId,
                title: title,
                content: content
            };

            console.log("Create note body: ", body);

            const response = await this.client.post('/createnote', body);

            // console.log("Create note response: ", response );

            return response.data;
        } catch (error) {
            console.error("Error creating note: ", error.message);
            throw error;
        }
    }

    async getNotes( userId ) {
        try {
            const body = {
                id: userId
              };

            const response = await this.client.get('/getnote', { params: body });

            console.log("Get notes response: ", response.data );

            return response.data;
        } catch (error) {
            console.error("Error getting notes: ", error.message);
            throw error;
        }
    }


    async updateNoteForUser(userId, noteId, title, content) {
        try {
          const body = {
            user_id: userId,
            id: noteId,
            title: title,
            content: content,
          };

          const response = await this.client.put("/updatenote", body);

          console.log("Update note response: ", response );

          return response;
        } catch (error) {
          console.error("Erreur lors de la mise à jour de la note :", error.message);
          throw error;
        }
      }

      async deleteNoteForUser(userId, noteId) {
        console.log("Deleteeeeeee note for user: ", noteId, userId);

        try {
          const body = {
            user_id: userId,
            id: noteId,
          };

          console.log("Delete note body: ", body);

          const response = await this.client.delete("/deletenote", { data: body });

          console.log("Delete note response: ", response );

          return response;
        } catch (error) {
          console.error("Erreur lors de la suppression de la note :", error.message);
          throw error;
        }
    }
}

const goServer = new GoServer();
export default goServer;
