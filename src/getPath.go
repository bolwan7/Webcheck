package main

import (
    "os"
    "path/filepath"
)

// Zwraca pełną ścieżkę do pliku data.txt znajdującego się w katalogu głównym projektu
func getDataFilePath() (string, error) {
    // Pobieramy bieżący katalog roboczy
    currentDir, err := os.Getwd()
    if err != nil {
        return "", err
    }

    // Szukamy ścieżki do katalogu głównego projektu
    projectDir := filepath.Dir(currentDir) // Wróć jeden poziom wyżej od `src`
    dataFilePath := filepath.Join(projectDir, "data.txt")

    return dataFilePath, nil
}