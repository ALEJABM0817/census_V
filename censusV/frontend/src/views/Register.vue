<template>
  <form @submit.prevent="register">
    <input type="email" v-model="username" placeholder="Username" />
    <input type="password" v-model="password" placeholder="Password" />
    <button type="submit">Register</button>
  </form>
</template>

<script>
import api from '../services/axiosConfig'; // Import the axios configuration

export default {
data() {
  return { username: '', password: '' };
},
methods: {
  async register() {
    try {
      const credentials = {
        username: this.username,
        password: this.password
      };
      const response = await api.post('/api/register', credentials, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      if (response.status === 200 || response.status === 201) {
        this.$router.push('/login');
      } else {
        console.error('Registration failed:', response.data.message);
      }
    } catch (error) {
      console.error('Error during registration:', error);
    }
  }
}
};
</script>
<style scoped>
form {
  display: flex;
  flex-direction: column;
  max-width: 300px;
  margin: 0 auto;
}

input {
  padding: 10px;
  margin: 10px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.error {
  color: red;
  font-size: 14px;
  margin-top: 10px;
}

p {
  text-align: center;
}

a {
  color: #007bff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style>