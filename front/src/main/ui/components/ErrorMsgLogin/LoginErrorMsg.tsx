import { Alert, AlertTitle, Snackbar } from '@mui/material'
import { useState } from 'react'

export type LoginErrorMsgPropsType = {
    /**
     * Customize message
     */
  ErrorMsg: string
}

export const LoginErrorMsg = ({ ErrorMsg }: LoginErrorMsgPropsType) => {
  const [isSnackOpen, setSnackOpen] = useState<boolean>(true)
  const handleClose = () => {
    setSnackOpen(false)
  }
  return (
    <Snackbar
      open={isSnackOpen}
      onClose={handleClose}
      autoHideDuration={4000}
      anchorOrigin={{ vertical: 'top', horizontal: 'left' }}
    >
      <Alert onClose={handleClose} variant="filled" severity="error" sx={{ width: '100%' }}>
        <AlertTitle>Error</AlertTitle>
        <strong>{ErrorMsg}</strong>
      </Alert>
    </Snackbar>
  )
}
