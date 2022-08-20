package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(ta *task) {
	t.tasks = append(t.tasks, ta)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	count := 0
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
		if tarea.completado == true {
			fmt.Println("Tarea completada")
		}

		if (len(tarea.descripcion) + 15) > count {
			count = len(tarea.descripcion) + 15
		}
		for i := 0; i < count; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) String() string {
	return fmt.Sprintf("Nombre: %s\nDescripcion: %s\nCompletado: %t", t.nombre, t.descripcion, t.completado)
}

func (t *task) marcarCompleto() {
	t.completado = true
}

func (t *task) actDescripcion(newDescripcion string) {
	t.descripcion = newDescripcion
}

func (t *task) actNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar el curso de Go",
		descripcion: "completar el curso de Go de Platzi en platzi day"}

	t2 := &task{
		nombre:      "Completar el curso avanzado de Go",
		descripcion: "completar el curso avanzado de Go de Platzi en platzi day"}

	t3 := &task{
		nombre:      "Comprar un plan de Platzi",
		descripcion: "Comprar un plan de platzi y seguir aprendiendo",
	}
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)
	t1.marcarCompleto()
	lista.imprimirLista()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Henrry"] = lista

	t4 := &task{
		nombre:      "Conseguir trabajo",
		descripcion: "Conseguir trabajo de Go para un Junior sin experiencia"}

	t5 := &task{
		nombre:      "Crear mi empresa",
		descripcion: "Crear mi eempresa para no depenmder de alguien",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Elibeth"] = lista2

	fmt.Println("Tareas de Henrry")
	mapaTareas["Henrry"].imprimirLista()

	fmt.Println("Tareas de Elibeth")
	mapaTareas["Elibeth"].imprimirLista()

}
