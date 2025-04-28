import React from 'react';
import { useForm } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import './ServiceAppointmentForm.css';

const schema = yup.object().shape({
  name: yup.string().required('Обязательное поле'),
  phone: yup.string().required('Обязательное поле').matches(/^\+?[0-9\s-]+$/, 'Некорректный номер'),
  vin: yup.string(),
  consent: yup.boolean().oneOf([true], 'Необходимо дать согласие'),
});

const ServiceAppointmentForm = () => {
  const { 
    register, 
    handleSubmit, 
    formState: { errors } 
  } = useForm({
    resolver: yupResolver(schema),
  });

  const onSubmit = (data) => {
    console.log('Форма отправлена:', data);
    alert('Форма успешно отправлена! С вами свяжутся в ближайшее время.');
  };

  return (
    <div className="service-appointment-form">
      <h1>ЗАПИСЬ НА ТЕХНИЧЕСКОЕ ОБСЛУЖИВАНИЕ</h1>
      <p className="form-description">
        Воспользуйтесь онлайн-формой для записи на техническое обслуживание, и специалист дилерского центра свяжется с вами в ближайшее время.
      </p>
      
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="form-row">
          <div className="form-field">
            <input
              id="name"
              type="text"
              placeholder="Имя *"
              {...register("name")}
              className={errors.name ? 'error' : ''}
            />
            {errors.name && <span className="error-message">{errors.name.message}</span>}
          </div>
          
          <div className="form-field">
            <input
              id="phone"
              type="tel"
              placeholder="Телефон *"
              {...register("phone")}
              className={errors.phone ? 'error' : ''}
            />
            {errors.phone && <span className="error-message">{errors.phone.message}</span>}
          </div>
        </div>
        
        <div className="form-field-long">
          <input
            id="vin"
            type="text"
            placeholder="VIN"
            {...register("vin")}
          />
        </div>
        
        <div className="form-notice">
          <p>* - поля, отмеченные звездочкой, обязательны к заполнению</p>
        </div>
        
        <div className="privacy-policy-container">
          <div className="privacy-policy-content">
            <p>1. Субъект персональных данных: лицо, обратившееся на сайт Toyota.ru в сети Интернет и предоставившее свои персональные данные Оператору (далее так же – «Пользователь»).</p>
            <p>2. Оператор обработки персональных данных: ООО «Тойота Мотор» (далее – «Тойота Мотор» или «Оператор»), имеющее место нахождения по адресу: 141031, Россия, Московская обл., Мытищинский р-н, МКАД, 84-й км, ПТЗ Аттурьевог, вл. 5, стр. 1.</p>
            <p>3. Цель обработки персональных данных: организация участия Пользователя в мероприятиях, в том числе в качестве участника мероприятия; предоставление информации Пользователю, включая информацию рекламного характера, в том числе, но не ограничиваясь, о наличии специальных предложений, проведении мероприятий, акций, презентаций в отношении продукции Toyota.</p>
          </div>
        </div>
        
        <div className="form-notice">Отправляя данную форму, Я даю согласие на обработку своих персональных данных.</div>
        
        
        <div className="marketing-consent">
          <input
            type="checkbox"
            id="marketingConsent"
            {...register("marketingConsent")}
          />
          <label htmlFor="marketingConsent">
            Настоящим в дополнение к целям, указанным в п. 1.3. Согласия, Я разрешаю обрабатывать свои персональные данные с целью продвижения товаров, работ, услуг Общества на рынке, путем осуществления прямых контактов, а также выражаю свое согласие на получение рекламной информации любым доступным способом, включая сети электросвязи, а также через сеть интернет.
          </label>
        </div>
        
        <button type="submit" className="submit-button">ОТПРАВИТЬ</button>
      </form>
    </div>
  );
};

export default ServiceAppointmentForm;