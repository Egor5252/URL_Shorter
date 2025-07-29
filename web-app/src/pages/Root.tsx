import { useEffect } from 'react';
import { Link, useNavigate  } from 'react-router-dom';
import './Root.css';

type SuccessResponse = {
  data: { message: string };
  error: null;
  success: true;
};

type ErrorResponse = {
  data: null;
  error: { message: string };
  success: false;
};

type ApiResponse = SuccessResponse | ErrorResponse;

function Root() {
    const navigate = useNavigate();

    const doLogin = async () => {
    try {
      const response = await fetch('/whoami', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      })

      const result: ApiResponse = await response.json()

      if (!result.success) {
        navigate('/login');
      }
    } catch (error) {
        navigate('/login');
    }
  }

  useEffect(() => {
    doLogin();
  }, []);

    return (
        <div>
            <div>
                <div className= "form-container">
                    <div className="hello-form">
                    <Link to="/changepass"><button>Сменить пароль</button></Link>
                    <Link to="/changeconfig"><button>Конфигурация системы</button></Link>
                    <Link to="/show"><button>Просмотр видео</button></Link>
                    <Link to="/logout"><button>Выйти</button></Link>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Root;

