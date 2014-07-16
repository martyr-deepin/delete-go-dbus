package transport

func (c *Conn) RemoveHandler(path ObjectPath) {
	c.handlersLck.Lock()
	delete(c.handlers, path)
	c.handlersLck.Unlock()
}
