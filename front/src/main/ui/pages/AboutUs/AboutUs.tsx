import img1 from '../../../../assets/img/AboutUs/Hotel.jpg';
import s from './AboutUs.module.css';

const {aboutUsTitle, aboutUsPhotoBlock, aboutUsMain, aboutText} = s;


export const AboutUs = () => {
    return (
        <div>
            <div className={aboutUsTitle}>
                <h1>About</h1>
            </div>
            <div className={aboutUsMain}>
                <div className={aboutText}><p>Our hotels, located just minutes from airport, offers resort style
                    boarding for your pets. Your best
                    friend will spend his or her vacation in loft style accommodations with an indoor pool and
                    playgrounds, and an outdoor splash park.</p>
                    <p>This is not the kind of place that crates Fido. Pet parents get their choice of having their dog
                        sleep in a shared lounge or in a private room. For around $200-a-night your furry companion can
                        even
                        stay in a poolside suite.</p>
                    <p>In addition to the outdoor pool, which is chemical-free, the 14,000-square-foot facility includes
                        a
                        playground, elevated orthopedic beds, filtered water, aromatherapy, and spa treatments. Yup,
                        your
                        dog will get better care than most people.</p>
                    <p>Oh and no need to install your own nanny-cam, the facility offers live webcams so you can keep
                        track of your pet from afar. The live feeds can be accessed from your computer or mobile phone
                        from 8am to 4pm every day of the week.</p>
                    <p>If the facility seems too nice to pass up, but you're not scheduled to take a trip, you can
                        always book your pet for their playcare service. Your dog can be dropped off for the day while
                        you work or run errands, like get a manicure or buy a yacht.
                    </p>
                    <p>Cat lovers need not fear, a separate wing is also available for cats.</p>
                </div>
                <div className={aboutUsPhotoBlock}>
                    <img src={img1} alt="hotelPhoto"/>
                </div>
            </div>

        </div>)
}