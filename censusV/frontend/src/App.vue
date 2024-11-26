<template>
  <div id="app">
    <header>
      <nav>
        <ul>
          <li><router-link to="/">Inicio</router-link></li>
          <li v-if="isAuthenticated"><router-link to="/census">Censos</router-link></li>
          <li v-if="!isAuthenticated"><router-link to="/login">Iniciar sesión</router-link></li>
          <li v-if="!isAuthenticated"><router-link to="/register">Registrarse</router-link></li>
          <li v-if="isAuthenticated" @click="logout">Cerrar sesión</li>
        </ul>
      </nav>
    </header>

    <main>
      <router-view />
    </main>

    <footer>
      <p>&copy; 2024 Mi Aplicación</p>
    </footer>
  </div>
</template>

<script>
import { ref, onMounted, provide } from 'vue';
import api from './services/axiosConfig';
import { useRouter } from 'vue-router';

export default {
  setup() {
    const isAuthenticated = ref(false);
    const authMessage = ref('');
    const router = useRouter();

    const checkAuthentication = async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        isAuthenticated.value = false;
        authMessage.value = 'No autorizado. Por favor, inicia sesión.';
        return;
      }

      try {
        const response = await api.get('/api/auth/check', {
          withCredentials: true,
          headers: {
            'Authorization': `Bearer ${token}`
          }
        });

        if (response.status === 200 && response.data) {
          isAuthenticated.value = response.data.isAuthenticated;
          authMessage.value = response.data.message || '';
        } else {
          isAuthenticated.value = false;
          authMessage.value = 'Error de autenticación. Intenta de nuevo.';
        }
      } catch (error) {
        if (error.response && error.response.status === 401) {
          authMessage.value = 'No autorizado. Por favor, inicia sesión.';
          router.push('/login');
        } else {
          isAuthenticated.value = false;
          authMessage.value = 'Error al verificar la autenticación. Por favor, verifica tu conexión.';
          console.error('Error checking authentication:', error);
        }
      }
    };

    const logout = async () => {
      try {
        await api.get('/api/logout', { withCredentials: true });
        localStorage.removeItem('token');
        isAuthenticated.value = false;
        authMessage.value = 'Has cerrado sesión exitosamente.';
        router.push('/login');
      } catch (error) {
        console.error('Error logging out:', error);
      }
    };

    onMounted(() => {
      checkAuthentication();
    });

    provide('checkAuthentication', checkAuthentication); 
    provide('authMessage', authMessage); 

    return {
      isAuthenticated,
      authMessage,
      logout
    };
  },
};
</script>

<style scoped>
#app {
  font-family: 'Arial', sans-serif;
  text-align: center;
}

header {
  background-color: #333;
  color: white;
  padding: 1rem;
}

nav ul {
  list-style-type: none;
  padding: 0;
}

nav ul li {
  display: inline;
  margin-right: 20px;
}

nav ul li a {
  color: white;
  text-decoration: none;
}

nav ul li a:hover {
  text-decoration: underline;
}

footer {
  margin-top: 2rem;
  padding: 1rem;
  background-color: #f4f4f4;
}
</style>

