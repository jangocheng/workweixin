package todos

func (m *ToDoResponse) MapResultList() []*ToDoList {
	rs := m.Result
	var rsp []*ToDoList
	for _, k := range rs {
		rsp = append(rsp, &ToDoList{
			ID:         k.ID,
			Name:       k.Name,
			CreateTime: k.CreateTime,
			FinishTime: k.FinishTime,
			Active:     k.Active,
		})
	}
	return rsp
}
