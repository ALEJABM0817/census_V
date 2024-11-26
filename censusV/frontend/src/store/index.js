const store = createStore({
    state: {
      user: null,
      token: localStorage.getItem('token') || '',
      preferences: {},
    },
    mutations: {
      setUser(state, user) {
        state.user = user;
      },
      setToken(state, token) {
        state.token = token;
        localStorage.setItem('token', token);
      },
      setPreferences(state, preferences) {
        state.preferences = preferences;
      },
      logout(state) {
        state.user = null;
        state.token = '';
        localStorage.removeItem('token');
      }
    },
    actions: {
      async login({ commit }, credentials) {
        try {
          const { data } = await axios.post('/api/login', credentials);
          commit('setUser', data.user);
          commit('setToken', data.token);
        } catch (error) {
          console.error("Login failed:", error);
        }
      },
      logout({ commit }) {
        commit('setToken', null);
        localStorage.removeItem('token');
      },
    

      async fetchPreferences({ commit, state }) {
        try {
          const { data } = await axios.get('/api/preferences', {
            headers: { Authorization: `Bearer ${state.token}` },
          });
          commit('setPreferences', data);
        } catch (error) {
          console.error("Failed to fetch preferences:", error);
        }
      },
    },
  });
  
  export default store;