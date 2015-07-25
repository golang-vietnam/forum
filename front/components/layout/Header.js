import React from 'react';

if (__CLIENT__) {
  var imgGo = require('../../images/go.png');
  console.log('imgGo', imgGo);
}

export default
class CounterButton extends React.Component {
  render() {
    return (
      <div className='ui main menu'>
        <div className='ui container'>
          <div href='#' className='header item'>
            <img className='logo'/>
            GolangVN
          </div>
          <a href='#' className='item'>Blog</a>
          <a href='#' className='item'>Articles</a>
          <a href='#' className='ui right floated dropdown item' tabIndex='0'>
            Dropdown <i className='dropdown icon'></i>
            <div className='menu' tabIndex='-1'>
              <div className='item'>Link Item</div>
              <div className='item'>Link Item</div>
              <div className='divider'></div>
              <div className='header'>Header Item</div>
              <div className='item'>
                <i className='dropdown icon'></i>
                Sub Menu
                <div className='menu'>
                  <div className='item'>Link Item</div>
                  <div className='item'>Link Item</div>
                </div>
              </div>
              <div className='item'>Link Item</div>
            </div>
          </a>
        </div>
      </div>
    );
  }
}
