<script setup>
import {
  PersonCircleOutline as PersonIcon,
} from "@vicons/ionicons5";
</script>
    
<template>
  <n-card>
    <n-row>
      <n-col :span="12">
        <n-statistic label="Users">
          <template #prefix>
            <n-icon>
              <person-icon />
            </n-icon>
          </template>
          1
        </n-statistic>
      </n-col>
    </n-row>
  </n-card>
  
    
</template>