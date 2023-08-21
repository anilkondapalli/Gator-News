import React, { useState, useCallback } from "react";
import { Form, Button, Row, Col, Spinner } from "react-bootstrap";
import DatePicker from "react-datepicker";
import es from "date-fns/locale/es";
import { useDropzone } from "react-dropzone";
import { API_HOST } from "../../../utils/constant";
import "./EditUserForm.scss";
import { toast } from "react-toastify";
import moment from "moment";
import {Camera} from "../../../utils/Icons"
import {
  uploadBannerApi,
  uploadAvatarApi,
  updateInfoApi,
} from "../../../api/user";

export default function EditUserForm(props) {
    const { user, setShowModal } = props;

    const [formData, setFormData] = useState(initialValue(user));
 
    const [bannerUrl, setBannerUrl] = useState(
      user?.banner ? `${API_HOST}/getBanner?id=${user.id}` : null
    );
    const [bannerFile, setBannerFile] = useState(null);

    const [avatarUrl, setAvatarUrl] = useState(
      user?.avatar ? `${API_HOST}/getAvatar?id=${user.id}` : null
    );
    const [avatarFile, setAvatarFile] = useState(null);
    const [loading, setLoading] = useState(false);

  const onDropBanner = useCallback((acceptedFile) => {
    const file = acceptedFile[0];
    setBannerUrl(URL.createObjectURL(file));
    setBannerFile(file);
  });

  const {
    getRootProps: getRootBannerProps,
    getInputProps: getInputBannerProps,
  } = useDropzone({
    accept: "image/jpeg, image/png",
    noKeyboard: true,
    multiple: false,
    onDrop: onDropBanner,
  });

    // eslint-disable-next-line react-hooks/exhaustive-deps
    const onDropAvatar = useCallback((acceptedFile) => {
      const file = acceptedFile[0];
      setAvatarUrl(URL.createObjectURL(file));
      setAvatarFile(file);
    });
    const {
      getRootProps: getRootAvatarProps,
      getInputProps: getInputAvatarProps,
    } = useDropzone({
      accept: "image/jpeg, image/png",
      noKeyboard: true,
      multiple: false,
      onDrop: onDropAvatar,
    });
    const onChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
      };

    const onSubmit = async(e) => {
      e.preventDefault();
      setLoading(true);


      if (bannerFile) {
        await uploadBannerApi(bannerFile).catch(() => {
          toast.error("Error uploading the cover picture");
        });
      }
      if (avatarFile) {
        await uploadAvatarApi(avatarFile).catch(() => {
          toast.error("Error uploading the profile picture");
        });
      }
      await updateInfoApi(formData)
      .then(() => {
        setShowModal(false);
      })
      .catch(() => {
        toast.error("Error al actualizar los datos");
      });

      setLoading(false);

      window.location.reload();
      
  };

  return (
    <div className="edit-user-form">
      <div
        className="banner"
        style={{ backgroundImage: `url('${bannerUrl}')` }}
        {...getRootBannerProps()}
      >
        <input {...getInputBannerProps()} />
        <Camera />
      </div>
      <div
        className="avatar"
        style={{ backgroundImage: `url('${avatarUrl}')` }}
        {...getRootAvatarProps()}>
        <input {...getInputAvatarProps()} />
        <Camera />
      </div>

      <Form onSubmit={onSubmit}>
        <Form.Group>
          <Row>
            <Col>
              <Form.Control type="text" placeholder="Name" name="name" defaultValue={formData.name} onChange={onChange} />
            </Col>
            <Col>
              <Form.Control type="text" placeholder="Last Name" name="lastName" defaultValue={formData.lastName} onChange={onChange} />
            </Col>
          </Row>
        </Form.Group>

        <Form.Group>
          <Form.Control
            as="textarea"
            row="3"
            placeholder="Add your biography"
            type="text"
            name="biography"
            defaultValue={formData.biography}
            onChange={onChange}
          />
        </Form.Group>

        <Form.Group>
          <Form.Control
            type="text"
            placeholder="Website"
            name="website"
            defaultValue={formData.website}
            onChange={onChange}

          />
        </Form.Group>

        <Form.Group>
          <DatePicker
            placeholder="Date of birth"
            locale={"en"}
            // dateFormat="MM-DD-YYYY"
            selected={new Date()}
            onChange={(value) =>
                setFormData({ ...formData, dateOfBirth: value })}
          />
        </Form.Group>

        <Button className="btn-submit" variant="primary" type="submit">
         {loading && <Spinner animation="border" size="sm" />} Update
        </Button>
      </Form>
    </div>
  );
}

function initialValue(user) {
    return {
      name: user.name || "",
      lastName: user.lastName || "",
      biography: user.biography || "",
      location: user.location || "",
      website: user.website || "",
    };
  }
