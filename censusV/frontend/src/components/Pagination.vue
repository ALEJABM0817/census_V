<template>
  <div class="pagination">
    <button @click="changePage(currentPage - 1)" :disabled="currentPage === 1">Anterior</button>
    <span>Página {{ currentPage }} de {{ totalPages }}</span>
    <button @click="changePage(currentPage + 1)" :disabled="currentPage === totalPages">Siguiente</button>
  </div>
</template>

<script>
export default {
  props: {
    totalPages: {
      type: Number,
      required: true
    },
    currentPage: {
      type: Number,
      required: true
    }
  },
  methods: {
    changePage(page) {      
      if (page > 0 && page <= this.totalPages) {
        this.$emit('page-change', page);
      }
    },
  },
  created() {
    console.log('Pagination component created', this.totalPages, this.currentPage);
  }
};
</script>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

button {
  padding: 10px 20px;
  background-color: #42a5f5;
  color: white;
  border: none;
  cursor: pointer;
}

button:disabled {
  background-color: #b0bec5;
  cursor: not-allowed;
}

span {
  margin: 0 10px;
}
</style>