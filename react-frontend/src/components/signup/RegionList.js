import React from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import Checkbox from '@material-ui/core/Checkbox'
import RegionsContext from '../../context/RegionsContext'

export default function RegionList({checked, setChecked}) {
  const regionsContext = React.useContext(RegionsContext)

  const handleToggle = (key) => () => {
    const currentIndex = checked.indexOf(key)
    const newChecked = [...checked]

    if (currentIndex === -1) {
      newChecked.push(key)
    } else {
      newChecked.splice(currentIndex, 1)
    }

    setChecked(newChecked)
  }

  return (
    <List dense>
      { regionsContext.regionsData.map((element) => {
        const key = element.key
        return (
          <ListItem onClick={handleToggle(key)} key={key} button>
            <Checkbox
              edge="end"
              onChange={handleToggle(key)}
              checked={checked.indexOf(key) !== -1}
              inputProps={{ 'aria-labelledby': key }}
            />
            &nbsp;&nbsp;&nbsp;&nbsp;<h1>{element.flag}</h1>&nbsp;&nbsp;&nbsp;&nbsp;<h5>{element.location}</h5>
          </ListItem>
        )
      })}
    </List>
  )
}