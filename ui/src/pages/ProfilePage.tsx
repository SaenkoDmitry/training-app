import React from "react";
import {useAuth} from "../context/AuthContext";
import Button from "../components/Button";

const ProfilePage: React.FC = () => {
    const {user, logout} = useAuth();

    if (!user) return null;

    return (
        <div
            style={{
                maxWidth: 420,
                margin: '0 auto',
                padding: '1rem',
                display: 'flex',
                flexDirection: 'column',
                gap: 20,
            }}
        >
            {/* карточка пользователя */}
            <div
                style={{
                    background: '#fff',
                    borderRadius: 20,
                    padding: '1.5rem',
                    boxShadow: '0 6px 20px rgba(0,0,0,0.06)',
                    textAlign: 'center',
                }}
            >
                <div style={{fontSize: 42, marginBottom: 8}}>👤</div>

                <div style={{fontSize: 18, fontWeight: 600}}>
                    {user.first_name}
                </div>

                {user.username && (
                    <div style={{opacity: 0.6, fontSize: 14}}>
                        @{user.username}
                    </div>
                )}
            </div>

            {/* logout */}
            <Button
                variant="danger"
                onClick={logout}
                style={{
                    width: '100%',
                    height: 48,
                    fontSize: 16,
                    borderRadius: 14,
                }}
            >
                Выйти из аккаунта
            </Button>
        </div>
    );
};

export default ProfilePage;
