import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';

// react的基础代码学习

// 最简单的React组件类
class Square extends React.Component {  // 定义一个自己的类，继承React的组件类
    render() {  // 写一个render方法, 表示他是怎么被渲染的
      return (  // 返回一个jsx   jsx被括号包裹 里面可以直接写HTML标签
        <button className="square">
          {this.props.val1}
        </button>
      );
    }
}

// 当Square类定义完了之后，我们就可以通过<Square />这个html标签来使用这个元素了。
// <Square /> 等同于 React.createElement('Square')
const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<Square />);  


class Board extends React.Component {
    renderSquare(i) {  
      return <Square val1={i} />;   // 传递数据的方法，在这里 父组件Board 通过prop把值 传递给了子组件 Square
    }

    render() {
      const status = 'Next player: X';
      return (
        <div>
          <div className="status">{status}</div>  // 在jsx中使用变量值
          <div className="board-row">
            {this.renderSquare(0)}   // 在jsx中可以直接调用本类的其他方法
            {this.renderSquare(1)}
          </div>
        </div>
      );
    }
}


// 通过state来在组件里记录状态
class Square extends React.Component {
    constructor(props) {  // 声明一个构造函数
      super(props); // super需要手动调
      this.state = {
        value: null,
      };
    }
    render() {
        return (  // 示例， state的Get和Set
          <button className="square"  onClick={() => this.setState({value: 'X'})} >
            {this.state.value}
          </button>
        );
      }
}


// 场景，当A和B两个组件之间需要相互通讯时 通常做法是把子组件包含的state数据提升至共同的父组件当中保存。
// 之后父组件拥有所以的数据，然后可以通过 props 将状态数据传递到子组件当中。
class Board extends React.Component {
    constructor(props) {
      super(props);
      this.state = { // Board作为一个父组件，直接保存9个square的状态，
        squares: Array(9).fill(null),
      };
    }
  
    renderSquare(i) {
      return <Square value={this.state.squares[i]} />;  // 这些square的状态通过props向下传递给子组件
    }
}