<template>
  <div class="filter-panel">
    <h2>Data Filters</h2>
    <div class="filter-group">
      <label for="education">Education:</label>
      <select id="education" v-model="filters.education">
        <option value="">All</option>
        <option value="Bachelors">Bachelors</option>
        <option value="Masters">Masters</option>
        <option value="Doctorate">Doctorate</option>
      </select>
    </div>
    <div class="filter-group">
      <label for="age">Age Range:</label>
      <select id="age" v-model="filters.age">
        <option value="">All</option>
        <option value="18-25">18-25</option>
        <option value="26-35">26-35</option>
        <option value="36-45">36-45</option>
        <option value="46-60">46-60</option>
        <option value="60+">60+</option>
      </select>
    </div>
    <div class="filter-group">
      <label for="income">Income Level:</label>
      <select id="income" v-model="filters.income">
        <option value="">All</option>
        <option value="<=50K">< $50,000</option>
        <option value=">50K">>= $50,000</option>
      </select>
    </div>
    <button @click="applyFilters">Apply Filters</button>
  </div>
</template>

<script>
export default {
  name: 'Filters',
  data() {
    return {
      filters: {
        education: '',
        age: '',
        income: '',
      },
    };
  },
  methods: {
    applyFilters() {
      this.$emit('apply-filters', this.filters);
    },
  },
};
</script>

<style scoped>
.filter-panel {
  padding: 1rem;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

.filter-group {
  margin-bottom: 1rem;
}

button {
  padding: 0.5rem 1rem;
  background-color: #42a5f5;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #1e88e5;
}
</style>