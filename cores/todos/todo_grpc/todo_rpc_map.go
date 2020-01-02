package todo_grpc

import "github.com/vnotes/workweixin_app/cores/todos"

func (m *ToDoResponse) MapResultList() []*todos.ToDoList {
	rs := m.Result
	var rsp []*todos.ToDoList
	for _, k := range rs {
		rsp = append(rsp, &todos.ToDoList{
			ID:         k.ID,
			Name:       k.Name,
			CreateTime: k.CreateTime,
			FinishTime: k.FinishTime,
			Active:     k.Active,
		})
	}
	return rsp
}
