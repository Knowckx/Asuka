import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';

// react的基础代码学习

// 最简单的React组件类
class Square extends React.Component {  // 定义一个自己的类，继承React的组件类
    render() {  // 写一个render方法, 表示他是怎么被渲染的
      return (  // 返回一个jsx   jsx被括号包裹 里面可以直接写HTML标签
        <button className="square">
          {/* TODO */}
        </button>
      );
    }
}

// 当Square类定义完了之后，我们就可以通过<Square />这个html标签来使用这个元素了。
// <Square /> 等同于 React.createElement('Square')
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<Square />);  


class Board extends React.Component {
    renderSquare(i) {   // 除了render方法，还可以定义其他的方法
      return <Square />;   
    }
  
    render() {
      const status = 'Next player: X';
      return (
        <div>
          <div className="status">{status}</div>  // 在jsx中使用变量值
          <div className="board-row">
            {this.renderSquare(0)}   // 如何调用本类的其他方法
            {this.renderSquare(1)}
          </div>
        </div>
      );
    }
}
  


  