import './assets/styles.css'

import { createApp } from 'vue';
import App from './App.vue';
import router from './router';

const app = createApp(App);

app.use(router);
app.mount('#app');

// Ensure App is properly set up before using its methods
const { checkAuthentication, authMessage } = App.setup();
app.provide('checkAuthentication', checkAuthentication); 
app.provide('authMessage', authMessage);
