<template>
  <div>
    <table>
      <thead>
        <tr>
          <th @click="sort('Age')">Age</th>
          <th @click="sort('Workclass')">Workclass</th>
          <th @click="sort('Education')">Education</th>
          <th @click="sort('MaritalStatus')">Marital Status</th>
          <th @click="sort('Occupation')">Occupation</th>
          <th @click="sort('Relationship')">Relationship</th>
          <th @click="sort('Race')">Race</th>
          <th @click="sort('Sex')">Sex</th>
          <th @click="sort('HoursPerWeek')">Hours Per Week</th>
          <th @click="sort('Income')">Income</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in data" :key="item.ID">
          <td>{{ item.Age }}</td>
          <td>{{ item.Workclass }}</td>
          <td>{{ item.Education }}</td>
          <td>{{ item.MaritalStatus }}</td>
          <td>{{ item.Occupation }}</td>
          <td>{{ item.Relationship }}</td>
          <td>{{ item.Race }}</td>
          <td>{{ item.Sex }}</td>
          <td>{{ item.HoursPerWeek }}</td>
          <td>{{ item.Income }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: ['data'],
  methods: {
    sort(attribute) {
      this.$emit('sort', attribute);
    },
  },
};
</script>