import * as React from 'react'
import Accordion from '@mui/material/Accordion'
import AccordionSummary from '@mui/material/AccordionSummary'
import AccordionDetails from '@mui/material/AccordionDetails'
import Typography from '@mui/material/Typography'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'
import Box from '@mui/material/Box'
import Card from '@mui/material/Card'
import CardActions from '@mui/material/CardActions'
import CardContent from '@mui/material/CardContent'
import Button from '@mui/material/Button'

const bull = (
  <Box component='span' sx={{ display: 'inline-block', mx: '2px', transform: 'scale(0.8)' }}>
    •
  </Box>
)

export default function ProjectCard() {
  return (
    <Card sx={{ maxWidth: 400, m: 5 }}>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color='text.secondary' gutterBottom>
          Project
        </Typography>
        <Typography sx={{ mb: 1.5 }} variant='h5' component='div'>
          ProjectName
        </Typography>
        <Typography sx={{ mb: 1.5 }} color='text.secondary'>
          ProjectDescriptionProjectDescription
        </Typography>
        <div>
          <Accordion>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls='panel1a-content'
              id='panel1a-header'
            >
              <Typography>Milestone1</Typography>
            </AccordionSummary>
            <AccordionDetails>
              <Button sx={{ width: '100%' }}>Task1</Button>
              <Button sx={{ width: '100%' }}>Task2</Button>
              <Button sx={{ width: '100%' }}>Task3</Button>
            </AccordionDetails>
          </Accordion>
          <Accordion>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls='panel2a-content'
              id='panel2a-header'
            >
              <Typography>Milestone2</Typography>
            </AccordionSummary>
            <AccordionDetails>
              <Button sx={{ width: '100%' }}>Task1</Button>
              <Button sx={{ width: '100%' }}>Task2</Button>
              <Button sx={{ width: '100%' }}>Task3</Button>
            </AccordionDetails>
          </Accordion>
          {/* <Accordion disabled>
            <AccordionSummary
              expandIcon={<ExpandMoreIcon />}
              aria-controls='panel3a-content'
              id='panel3a-header'
            >
              <Typography>Disabled Accordion</Typography>
            </AccordionSummary>
          </Accordion> */}
        </div>
      </CardContent>
      <CardActions>
        <Button size='small'>Learn More</Button>
      </CardActions>
    </Card>
  )
}
