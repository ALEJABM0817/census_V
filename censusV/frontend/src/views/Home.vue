<template>
  <div>
    <h1>Datos del Censo</h1>
    
    <!-- Filtros -->
    <Filters @apply-filters="updateFilters" />
    
    <!-- Tabla de Datos -->
    <DataTable :data="data" @sort="fetchData" />
    
    <!-- Paginación -->
    <Pagination @page-change="changePage" :totalPages="totalPages" :currentPage="currentPage" />
  </div>
</template>

<script>
import axios from 'axios';
import Filters from '@/components/Filters.vue';
import DataTable from '@/components/DataTable.vue';
import Pagination from '@/components/Pagination.vue';

export default {
  components: { Filters, DataTable, Pagination },
  data() {
    return {
      data: [],
      totalPages: 0,
      currentPage: 1,
      filters: {
        education: '',
        income: '',
        occupation: '',
        age: ''
      },
    };
  },
  methods: {
    async fetchData() {
      const params = {
        page: this.currentPage,
        ...this.filters,
      };
      try {
        const response = await axios.get('/api/census', { params });
        this.data = response.data.items;
        this.totalPages = response.data.totalPages;
        if (response.data.currentPage) {
          this.currentPage = response.data.currentPage; // Ensure currentPage is updated
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    updateFilters(filters) {
      this.filters = { ...this.filters, ...filters };
      this.currentPage = 1; // Reset to first page when filters are updated
      this.fetchData();
    },
    changePage(page) {
      this.currentPage = page;
      this.fetchData();
    },
  },
  created() {
    this.fetchData();
  },
};
</script>
  
  <style scoped>
  .home {
    text-align: center;
    padding: 20px;
  }
  
  h1 {
    font-size: 2rem;
    margin-bottom: 20px;
  }
  
  .filters {
    margin-bottom: 20px;
  }
  
  .statistics {
    margin-top: 20px;
  }
  
  .export {
    margin-top: 20px;
  }
  
  button {
    padding: 10px 20px;
    background-color: #28a745;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  button:hover {
    background-color: #218838;
  }
  </style>