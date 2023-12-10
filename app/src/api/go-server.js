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
}

const goServer = new GoServer();
export default goServer;
