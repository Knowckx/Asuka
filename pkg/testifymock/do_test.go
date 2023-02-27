package testifymock

import (
	"fmt"
	"testing"

	infa "github.com/Knowckx/ainfa"
	"github.com/stretchr/testify/mock"
)

// 1.新建一个结构体，包含mock.Mock对象
type MockCrawler struct {
	mock.Mock
}

// 2.这个结构体 写一个和目标一样的方法
func (m *MockCrawler) GetUserList() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

func TestGetUserList(t *testing.T) {
	MockUsers := []*User{}
	MockUsers = append(MockUsers, &User{"dj", 18})
	MockUsers = append(MockUsers, &User{"zhangsan", 20}) // 3. 构造返回值
	crawler := new(MockCrawler)                          // 4. new一个mock的结构体
	crawler.On("GetUserList").Return(MockUsers, nil)     // 5. 通过on来实现对应的方法的打桩

	res, err := crawler.GetUserList()
	infa.PrintJson(res)
	fmt.Println(err)
	// 实际应用中，把crawler赋值给对应的接口对象就好了
}
