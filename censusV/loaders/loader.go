package loaders

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"

    "censusV/database"
    "censusV/models"
)

// LoadCensusData carga datos censales desde un archivo y los guarda en la base de datos.
func LoadCensusData(filePath string) error {
    // Abrir el archivo
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("error al abrir el archivo: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var wg sync.WaitGroup
    db := database.DB

    // Leer la primera línea del archivo
    if scanner.Scan() {
        firstLine := scanner.Text()
        fmt.Println("Primera línea del archivo:", firstLine)
    }

    // Leer cada línea del archivo
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Split(line, ",")
        if len(fields) < 10 {
            // Ignorar líneas con campos incompletos
            fmt.Printf("Línea ignorada por datos incompletos: %s\n", line)
            continue
        }

        // Convertir los valores necesarios
        age, err := strconv.Atoi(fields[0])
        if err != nil {
            fmt.Printf("Error al convertir la edad: %v para la línea: %s\n", err, line)
            continue
        }

        hours, err := strconv.Atoi(fields[8])
        if err != nil {
            fmt.Printf("Error al convertir las horas: %v para la línea: %s\n", err, line)
            continue
        }

        // Crear una estructura CensusData
        data := models.CensusData{
            Age:           age,
            Workclass:     strings.TrimSpace(fields[1]),
            Education:     strings.TrimSpace(fields[2]),
            MaritalStatus: strings.TrimSpace(fields[3]),
            Occupation:    strings.TrimSpace(fields[4]),
            Relationship:  strings.TrimSpace(fields[5]),
            Race:          strings.TrimSpace(fields[6]),
            Sex:           strings.TrimSpace(fields[7]),
            HoursPerWeek:  hours,
            Income:        strings.TrimSpace(fields[9]),
        }

        // Usar una goroutine para guardar datos concurrentemente
        wg.Add(1)
        go func(data models.CensusData) {
            defer wg.Done()
            if err := db.Create(&data).Error; err != nil {
                fmt.Printf("Error al guardar los datos en la base: %v\n", err)
            }
        }(data)
    }

    // Esperar a que todas las goroutines terminen
    wg.Wait()

    // Manejar errores de escaneo
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error al leer el archivo: %v", err)
    }

    fmt.Println("Carga de datos completada exitosamente.")
    return nil
}