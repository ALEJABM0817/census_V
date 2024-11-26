<template>
  <form @submit.prevent="login">
    <input type="username" v-model="username" placeholder="Username" />
    <input type="password" v-model="password" placeholder="Password" />
    <button type="submit">Login</button>
  </form>
</template>

<script>
import api from '../services/axiosConfig';
import { useRouter } from 'vue-router';
import { inject } from 'vue';

export default {
  data() {
    return {
      username: '',
      password: ''
    };
  },
  setup() {
    const router = useRouter();
    const checkAuthentication = inject('checkAuthentication'); 

    return {
      checkAuthentication,
      router
    };
  },
  methods: {
    async login() {
      try {
        const credentials = {
          username: this.username,
          password: this.password
        };
        const response = await api.post('/api/login', credentials, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json'
          }
        });
        if (response.status === 200) {
          localStorage.setItem('token', response.data.token);
          this.$emit('authenticated');
          this.checkAuthentication(); 
          this.$router.push('/home');
        } else {
          console.error('Login failed:', response.data.message);
        }
      } catch (error) {
        console.error('Error during login:', error);
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