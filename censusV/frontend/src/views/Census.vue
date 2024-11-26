<template>
  <div>
    <h1>Census Data</h1>
    <!-- Filters -->
    <Filters @apply-filters="updateFilters" />
    <!-- Data Table -->
    <DataTable :data="data" @sort="sortData" />
    <!-- Pagination -->
    <Pagination @page-change="changePage" :totalPages="totalPages" :currentPage="currentPage" />
  </div>
</template>

<script>
import api from '../services/axiosConfig';
import Filters from "@/components/Filters.vue";
import DataTable from "@/components/DataTable.vue";
import Pagination from "@/components/Pagination.vue";

export default {
  components: { Filters, DataTable, Pagination },
  data() {
    return {
      data: [],
      totalPages: 0,
      currentPage: 1,
      filters: {
        education: "",
        income: "",
        occupation: "",
        age: "",
      },
      sortAttribute: '',
      sortOrder: 'asc',
    };
  },
  methods: {
    async fetchData() {
      const params = {
        page: this.currentPage,
        limit: 2,
        offset: (this.currentPage - 1) * 2,
        education: this.filters.education,
        income: this.filters.income,
        occupation: this.filters.occupation,
        age: this.filters.age,
        order_by: this.sortAttribute,
        direction: this.sortOrder,
      };
      try {
        const response = await api.get('/api/census', { 
          params,
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
          }
        });
        if (response.status === 200) {
          console.info('Data fetched:', response.data);
          this.data = response.data.paginated_data; // Update data with fetched data
          // Assuming the response includes total pages
          this.totalPages = Math.ceil(response.data.records / params.limit) // Math.ceil(response.data.length / 10); // Update total pages
        } else {
          console.error('Login failed:', response.data.message);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    updateFilters(filters) {
      this.filters = { ...this.filters, ...filters };
      this.fetchData();
    },
    changePage(page) {
      this.currentPage = page;
      this.fetchData();
    },
    sortData(attribute) {
      this.sortAttribute = attribute;
      this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc';
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
