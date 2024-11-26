<template>
  <form @submit.prevent="register">
    <input type="email" v-model="username" placeholder="Username" />
    <input type="password" v-model="password" placeholder="Password" />
    <button type="submit">Register</button>
  </form>
</template>

<script>
export default {
data() {
  return { username: '', password: '' };
},
methods: {
  async register() {
    try {
      await fetch('http://localhost:8080/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username: this.username, password: this.password })
      });
      this.$router.push('/login');
    } catch (error) {
      console.error('Error during registration:', error);
    }
  }
}
};
</script>