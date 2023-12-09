import axios from 'axios';

class GoServer {
  constructor() {
    this.client = axios.create({
      baseURL: 'http://localhost:8080',
    });
  }

    async signUp({ username, email, password }) {
        try {
            const body = {
                username: username,
                email: email,
                password: password
            };

            const response = await this.client.post('/signup', body);
            return response.data;
        } catch (error) {
            console.error("Error signing up: ", error.message);
            throw error;
        }
    }

    async logIn({ email, password }) {
        try {
            const body = {
                email: email,
                password: password
            };

            const response = await this.client.post('/login', body);
            return response.data;
        } catch (error) {
            console.error("Error signing in: ", error.message);
            throw error;
        }
    }

    async getUsers() {
        try {
            const response = await this.client.get('/user/getall');
            return response.data;
        } catch (error) {
            console.error("Error getting users: ", error.message);
            throw error;
        }
    }
}

const goServer = new GoServer();
export default goServer;
