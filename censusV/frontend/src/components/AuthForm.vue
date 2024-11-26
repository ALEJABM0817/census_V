<template>
  <div>
    <form @submit.prevent="submitForm">
      <input type="email" v-model="email" placeholder="Email" required />
      <input type="password" v-model="password" placeholder="Password" required />
      <button type="submit">{{ isLogin ? 'Login' : 'Register' }}</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    isLogin: { type: Boolean, default: true },
  },
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    async submitForm() {
      const url = this.isLogin ? '/api/login' : '/api/register';
      try {
        const response = await axios.post(url, {
          email: this.email,
          password: this.password,
        });
        this.$store.commit('setToken', response.data.token);
        this.$router.push('/');
      } catch (error) {
        alert('Error: ' + error.response.data.message);
      }
    },
  },
};
</script>